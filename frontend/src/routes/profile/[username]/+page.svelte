<script lang="ts">
  export let data: { user: any; posts: any[] };

  let { user, posts } = data;
</script>

<div>
  {#if user}
    <h2>@{user.username}</h2>
    {#each posts as post}
      <div>
        {#if post.photo_url}
          <img src={post.photo_url} alt="post image" />
        {/if}
        <p>{post.description}</p>
      </div>
    {/each}
  {:else}
    <p>User not found</p>
  {/if}
</div>
