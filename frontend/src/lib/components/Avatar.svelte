<script lang="ts">
  export let src: string | undefined = undefined
  export let alt: string | undefined = ''
  export let size: number = 40
</script>

{#if src}
  <img src={src} alt={alt} class="rounded-full object-cover" style="width: {size}px; height: {size}px;" loading="lazy" decoding="async" />
{:else}
  <div class="rounded-full bg-gradient-to-br from-blue-400 to-purple-500 flex items-center justify-center text-white font-semibold" style="width: {size}px; height: {size}px;">{(alt && alt[0]) ? alt[0].toUpperCase() : '?'}</div>
{/if}
