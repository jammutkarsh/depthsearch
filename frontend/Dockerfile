# Choose a suitable Node.js base image:
FROM node:18-alpine

WORKDIR /app

# Copy package.json to install dependencies later
COPY package.json ./

# Install dependencies
RUN npm install

# Copy the rest of the project files
COPY . .

# Install dependencies
RUN npm run build

# Start the preview server:
CMD ["npm", "start"]
