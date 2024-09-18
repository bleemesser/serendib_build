#!/bin/bash

# Install Git and Nginx
sudo apt update
sudo apt upgrade -y
sudo apt install -y nginx python-certbot-nginx
sudo snap install --classic certbot
sudo ln -s /snap/bin/certbot /usr/bin/certbot

# Copy nginx.conf to /etc/nginx/nginx.conf and reload Nginx
sudo useradd nginx
sudo cp nginx.conf /etc/nginx/nginx.conf
sudo systemctl restart nginx
sudo nginx -s reload

sudo certbot --nginx -d survey.projectserendib.com

./main serve --http "127.0.0.1:8090"