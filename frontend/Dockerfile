# Build stage
FROM node:20-alpine AS builder

WORKDIR /app

# Копируем package.json
COPY package*.json ./
RUN npm ci

# Копируем исходный код
COPY . .

# Аргумент для API URL
ARG VITE_API_URL
ENV VITE_API_URL=$VITE_API_URL

# Собираем приложение
RUN npm run build

# Production stage
FROM nginx:alpine

# Копируем собранное приложение
COPY --from=builder /app/build /usr/share/nginx/html

# Копируем конфиг nginx для SPA
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]