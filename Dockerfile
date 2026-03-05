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

# Set the backend API URL for pre-rendering blog content at build time
ARG NUXT_API_BASE_URL=https://ivmanto.com
ENV NUXT_API_BASE_URL=${NUXT_API_BASE_URL}
ENV BACKEND_URL=${NUXT_API_BASE_URL}

# Generate the static site (pre-renders all routes defined in routeRules)
RUN npx nuxi generate

# Stage 2: Create the final, lightweight production image
FROM node:20-bookworm-slim

# Apply the latest available security patches to the final image.
RUN apt-get update && apt-get upgrade -y && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Install a production-ready static web server
RUN npm install -g serve

WORKDIR /app

# Copy the generated static files from the build stage.
# Nuxt SSG outputs to .output/public (not dist).
COPY --from=build /app/.output/public .

# Expose the port Cloud Run will listen on. 'serve' automatically uses the PORT env var.
EXPOSE 8080

# Start the static file server.
# The '-s' flag rewrites missing routes to index.html, which is needed for
# client-rendered pages (blog, booking) that don't have pre-rendered HTML.
CMD ["serve", "-s", "."]
