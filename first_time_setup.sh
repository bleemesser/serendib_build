#!/bin/bash

# Install Git and Nginx
sudo apt update
sudo apt upgrade -y
sudo apt install -y nginx

# Copy nginx.conf to /etc/nginx/nginx.conf and reload Nginx
sudo useradd nginx
sudo cp nginx.conf /etc/nginx/nginx.conf
sudo systemctl restart nginx
sudo nginx -s reload

# Create a systemd service file to run the server
cat << EOF | sudo tee /etc/systemd/system/pocketbase.service > /dev/null
[Unit]
Description=Pocketbase Service
After=network.target

[Service]
User=$USER
Group=$(id -gn)
WorkingDirectory=$(pwd)
ExecStart=$(pwd)/main serve --http "127.0.0.1:8090"
Restart=unless-stopped

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd
sudo systemctl daemon-reload

# Enable and start the service
sudo systemctl enable pocketbase.service
sudo systemctl restart pocketbase.service