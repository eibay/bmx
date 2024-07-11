# Use the official Node.js image
FROM node:16-alpine as build

# Set the working directory
WORKDIR /ui

# Copy the package.json and install dependencies
COPY package.json ./
RUN npm install

# Copy the rest of the application
COPY . .

# Build the application
RUN npm run build

# Use the official nginx image as the base image
FROM nginx:alpine

# Copy the build output to the nginx html directory
COPY --from=build /ui/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Command to run nginx
CMD ["nginx", "-g", "daemon off;"]
