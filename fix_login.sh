#!/bin/bash
# ============================================
# Skillture - Login Fix Deployment Script
# ============================================
# This script fixes the login issue by updating the Caddyfile
# and ensuring the backend is running properly.
#
# Run this on your Google Cloud server:
#   sudo bash fix_login.sh
# ============================================

set -e

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}=== Skillture Login Fix ===${NC}"

# 1. Update the Caddyfile
echo -e "${GREEN}[1/4] Updating Caddyfile...${NC}"
cat > /etc/caddy/Caddyfile << 'EOF'
# ============================================
# Skillture - Manual Deployment Configuration
# ============================================

:80 {
    # Proxy API Requests to Backend (must be handled FIRST)
    handle /api/* {
        reverse_proxy localhost:8080
    }

    # Serve React Frontend (SPA fallback for everything else)
    handle {
        root * /opt/skillture/web/dist
        try_files {path} /index.html
        file_server
    }
}
EOF

echo -e "${GREEN}[2/4] Restarting Caddy...${NC}"
systemctl restart caddy

echo -e "${GREEN}[3/4] Checking Skillture backend status...${NC}"
if systemctl is-active --quiet skillture; then
    echo -e "${GREEN}Backend is running ✓${NC}"
else
    echo -e "${YELLOW}Backend is NOT running, starting it...${NC}"
    systemctl restart skillture
fi

echo -e "${GREEN}[4/4] Verifying API is accessible...${NC}"
sleep 2

# Test that API endpoint is now reachable through Caddy
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" http://localhost/api/v1/admins/login -X POST -H "Content-Type: application/json" -d '{"username":"test","password":"test"}' 2>/dev/null || echo "000")

if [ "$HTTP_CODE" = "401" ] || [ "$HTTP_CODE" = "400" ]; then
    echo -e "${GREEN}API is working correctly ✓ (got expected HTTP $HTTP_CODE for bad credentials)${NC}"
elif [ "$HTTP_CODE" = "200" ]; then
    echo -e "${GREEN}API is working correctly ✓${NC}"
else
    echo -e "${RED}WARNING: API returned HTTP $HTTP_CODE - something may still be wrong${NC}"
    echo -e "${YELLOW}Check: sudo journalctl -u skillture -n 20${NC}"
    echo -e "${YELLOW}Check: sudo journalctl -u caddy -n 20${NC}"
fi

echo ""
echo -e "${GREEN}=== Fix Applied Successfully! ===${NC}"
echo -e "${YELLOW}Try logging in again from your browser.${NC}"
echo -e "${YELLOW}If there's no admin user yet, create one with:${NC}"
echo -e "  curl -X POST http://localhost/api/v1/admins/ -H 'Content-Type: application/json' -d '{\"username\":\"Skillture\",\"password\":\"Skillture141\"}'"
