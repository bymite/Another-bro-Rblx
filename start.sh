#!/bin/bash
# Auto-detect the primary network interface
IFACE=$(ip route | awk '/default/ {print $5; exit}')
echo "Detected interface: $IFACE"

# Patch the config with the real interface name
sed -i "s/IFACE/$IFACE/g" /etc/sockd.conf

echo "Starting dante with config:"
cat /etc/sockd.conf

exec sockd -f /etc/sockd.conf
