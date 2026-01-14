import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
  const username = params.username;

  const base = import.meta.env.VITE_API_URL || 'http://localhost:8080';

  // Use absolute backend URL for SSR so server-side rendering fetches the API directly
  const userRes = await fetch(`${base}/users/${username}`);
  const user = userRes.ok ? await userRes.json() : null;

  const postsRes = await fetch(`${base}/posts/users/${username}`);
  const posts = postsRes.ok ? await postsRes.json() : [];

  return { user, posts };
};
