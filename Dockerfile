# Stage 1: Build the application
# Use the latest LTS Node.js 20 runtime based on Debian 12 (Bookworm)
FROM node:20-bookworm-slim AS build

# Apply the latest available security patches from the OS vendor.
RUN apt-get update && apt-get upgrade -y && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Set the working directory in the container
WORKDIR /app

# Copy package.json and lock file to leverage Docker layer caching
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application's source code
COPY . .

# Build the Nuxt application for production
RUN npm run build

# Stage 2: Create the final, lightweight production image
FROM node:20-bookworm-slim

# Apply the latest available security patches to the final image.
RUN apt-get update && apt-get upgrade -y && \
    apt-get clean && rm -rf /var/lib/apt/lists/*
WORKDIR /app

# Copy the standalone, built application from the 'build' stage
COPY --from=build /app/dist .

# Expose the port the app will run on and start the server
CMD ["node", "server/index.mjs"]
