# Use the official Node.js image as the base
FROM node:16

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package.json pnpm-lock.yaml* ./

# Install PNPM globally
RUN npm install -g pnpm

# Install project dependencies using PNPM
RUN pnpm install

# Copy the rest of the application code to the working directory
COPY . .

# Expose the port your Node.js application listens on
EXPOSE 8080

# Define the command to run your Node.js application
CMD ["pnpm", "start"]
