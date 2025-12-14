export async function getUserPosts(user_id: string) {
  const res = await fetch(`http://localhost:8080/api/posts/${user_id}`);

  if (!res.ok) {
    throw new Error("Failed to load posts");
  }

  return res.json(); 
}
