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

# Install dependencies using the lockfile for a deterministic build
RUN npm ci

# Copy the rest of the application's source code
COPY . .

# Build the Nuxt application for production
RUN npm run build

# Stage 2: Create the final, lightweight production image
FROM node:20-bookworm-slim
 
# Node.js apps run as the 'node' user by default, which is a security best practice.
# Set the working directory.
WORKDIR /app
 
# Copy package manifests from the build stage.
COPY --from=build /app/package.json ./package.json
COPY --from=build /app/package-lock.json ./package-lock.json

# Install ONLY production dependencies. This creates a smaller and more secure image.
RUN npm ci --omit=dev
 
# Copy the built application assets
COPY --from=build /app/dist .
 
# Expose the port Cloud Run will listen on. The 'serve' package automatically
# respects the PORT environment variable provided by Cloud Run.
EXPOSE 8080
 
# Use the 'start' script from package.json to run the server. This ensures
# the execution is consistent and dependencies are managed locally.
CMD ["npm", "start"]
