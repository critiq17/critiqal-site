import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
  const username = params.username;

  const userRes = await fetch(`http://localhost:8080/api/users/${username}`);
  const user = userRes.ok ? await userRes.json() : null;

  const postsRes = await fetch(`http://localhost:8080/api/posts/users/${username}`);
  const posts = postsRes.ok ? await postsRes.json() : [];

  return { user, posts };
};
