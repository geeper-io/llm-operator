FROM node:18-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy source code
COPY . .

# Expose the dev server port
EXPOSE 3000

# Start the dev server
CMD ["npm", "start", "--", "--host", "0.0.0.0"]
