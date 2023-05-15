# Use the official nginx image as base image.
FROM nginx

# Copy the nginx configuration file.
COPY nginx.conf /etc/nginx/nginx.conf

# Copy our static files into the nginx public folder.
WORKDIR /usr/share/nginx/html

# Copy the static files from the build stage.
COPY --from=0 /app/build .

# Expose port 80 to the Docker host, so we can access it
EXPOSE 80

# Run nginx in the foreground.
CMD ["nginx", "-g", "daemon off;"]
