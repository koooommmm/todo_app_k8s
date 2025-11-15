# ---- Build stage ----
FROM node:lts-alpine3.22 AS builder

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install

COPY . .
RUN npm run build

# ---- Serve stage ----
FROM nginx:mainline

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
