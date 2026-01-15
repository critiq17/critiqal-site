# Critiqal Frontend

A **production-ready, enterprise-grade Svelte frontend** for the Critiqal social platform with Instagram-style authentication, advanced state management, and comprehensive UI component library.

## ✨ Key Features

- ✅ **100% TypeScript** - Full type safety with strict mode
- 🔐 **Cookie-Based Authentication** - Secure httpOnly cookies with token refresh
- 📱 **Mobile-First Design** - Responsive Tailwind CSS architecture
- 🎨 **Component Library** - Reusable, accessible, production-ready components
- 🔄 **Advanced State Management** - Svelte stores with derived computed properties
- ⚡ **Performance Optimized** - Code splitting, lazy loading, minification
- ♿ **Accessibility** - WCAG AA compliant with ARIA labels
- 🌙 **Dark Mode Support** - Theme switching with persistence
- 📡 **Type-Safe API Client** - Automatic error handling and token refresh
- 🚨 **Error Boundaries** - Comprehensive error handling everywhere

## 🚀 Quick Start

### Prerequisites
- Node.js 18+
- npm or pnpm

### Installation
```bash
npm install
```

### Development
```bash
npm run dev
```

Open [http://localhost:5173](http://localhost:5173)

### Production Build
```bash
npm run build
npm run preview
```

## 📁 Project Structure

See [ARCHITECTURE.md](./ARCHITECTURE.md) for complete directory structure and detailed documentation.

```
src/
├── lib/
│   ├── components/        # Reusable UI components
│   ├── stores/           # State management
│   ├── services/         # API clients
│   ├── types/            # TypeScript interfaces
│   ├── utils/            # Helpers & utilities
│   ├── hooks/            # Custom Svelte hooks
│   └── feature/          # Feature modules
└── routes/               # SvelteKit pages
```

## 🏗️ Architecture Highlights

### Type-Safe Components
```svelte
<script lang="ts">
  import Button from '$lib/components/Button.svelte'
  import type { SvelteHTMLElements } from 'svelte/elements'
  
  interface $$Props {
    variant?: 'primary' | 'secondary'
    size?: 'sm' | 'md' | 'lg'
    isLoading?: boolean
  }
</script>
```

### Reactive Stores
```typescript
import { user, isAuthenticated } from '$lib/stores/auth'

export function setUser(user: User) {
  authStore.update(state => ({
    ...state,
    user,
    isAuthenticated: true
  }))
}
```

### Type-Safe API Client
```typescript
import { get, post } from '$lib/services/api'

// Automatic error handling, token refresh, CSRF protection
const user = await post<User>('/auth/sign-in', credentials)
```

## 🔐 Cookie-Based Authentication

### Frontend
- httpOnly cookies for session storage
- Automatic token refresh on 401
- CSRF protection
- Persistent session state

### Backend Requirements
Backend should:
1. Set httpOnly, Secure, SameSite cookies
2. Implement `/auth/refresh` endpoint
3. Handle CSRF validation
4. Return proper 401 on token expiration

[Read Authentication Guide](./ARCHITECTURE.md#-cookie-security-configuration)

## 🧩 Component Library

### Button Component
```svelte
<Button variant="primary" size="md" on:click={handleClick}>
  Click me
</Button>
```

### Input Component
```svelte
<Input 
  type="email" 
  label="Email" 
  error={errors.email}
  required
/>
```

### Card Component
```svelte
<Card padding="lg">
  <h2>Card Title</h2>
  <p>Card content</p>
</Card>
```

### Toast Notifications
```typescript
import { notifications } from '$lib/stores/notifications'

notifications.success('Operation successful!')
notifications.error('Something went wrong')
notifications.info('Just FYI')
notifications.warning('Be careful!')
```

## 📋 Scripts

| Script | Description |
|--------|-------------|
| `npm run dev` | Start development server |
| `npm run build` | Build for production |
| `npm run preview` | Preview production build |
| `npm run lint` | Run ESLint |
| `npm run format` | Format code with Prettier |
| `npm run check` | Type check with TypeScript |
| `npm run check:watch` | Type check in watch mode |

## 🔒 Security Best Practices

✅ **Implemented**
- httpOnly cookies (no JS access)
- Secure flag (HTTPS only in production)
- SameSite protection (CSRF)
- CORS configured for credentials
- Token refresh on 401
- No sensitive data in localStorage
- Input validation on all forms
- Error boundaries prevent crashes

## 📱 Responsive Design

Mobile-first approach using Tailwind breakpoints:
```svelte
<div class="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
  <!-- Responsive grid -->
</div>
```

## ♿ Accessibility

- Semantic HTML elements
- ARIA labels on interactive elements
- Keyboard navigation support
- Focus visible indicators
- Color contrast WCAG AA compliant
- Error messages linked to inputs

## 🎨 Theming

### Dark Mode
```svelte
<script>
  import { theme } from '$lib/stores/theme'
  
  theme.setMode('dark')
</script>
```

### Custom Colors
Update `tailwind.config.js`:
```javascript
module.exports = {
  theme: {
    extend: {
      colors: {
        primary: '#3B82F6'
      }
    }
  }
}
```

## 🚀 Deployment

### Vercel (Recommended)
```bash
npm run build
git push  # Auto-deploys
```

### Docker
```dockerfile
FROM node:18-alpine
WORKDIR /app
COPY . .
RUN npm install && npm run build
EXPOSE 3000
CMD ["node", "-e", "require('./build/index.js').default({listen: {port: 3000}})"]
```

### Environment Variables
Set in your hosting platform:
```
VITE_API_URL=https://api.critiqal.com/api
```

## 📊 Performance

- **Code Splitting**: Automatic per route
- **Lazy Loading**: Images with WebP/AVIF
- **Minification**: JavaScript and CSS minified
- **Tree Shaking**: Unused code removed
- **Bundle Size**: ~50KB gzipped

## 🧪 Code Quality

All files pass:
- TypeScript strict mode
- ESLint rules
- Prettier formatting
- Accessibility checks

## 📖 Documentation

- [Architecture Guide](./ARCHITECTURE.md) - Detailed technical documentation
- [Component Library](./src/lib/components/README.md) - Component documentation
- [API Services](./src/lib/services/README.md) - Service documentation
- [Type Definitions](./src/lib/types/index.ts) - All TypeScript types

## 🤝 Contributing

1. Follow TypeScript strict mode
2. Use proper component props interfaces
3. Add error boundaries
4. Include loading states
5. Test mobile responsiveness
6. Run `npm run format && npm run lint`

## 📝 License

Proprietary - Critiqal Platform

## 🔗 Links

- [Backend Repository](../backend)
- [Main Website](https://critiqal.com)
- [Documentation](./ARCHITECTURE.md)

---

**Built with ❤️ using Svelte 5 + SvelteKit 2 + Tailwind CSS 4**
