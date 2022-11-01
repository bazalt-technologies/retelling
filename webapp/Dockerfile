FROM node:lts-alpine

# make the 'app' folder the current working directory
WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY package*.json ./

# install project dependencies
RUN npm install --force

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

# build app for production with minification
CMD ["npm", "i", "vuex@3"]
CMD ["npm", "run", "serve", "--", "--port", "8082"]