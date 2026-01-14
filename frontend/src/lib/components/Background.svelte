<script lang="ts">
  // Decorative background - blurred blobs + subtle palms
  // Purely visual; pointer-events: none
</script>

<div class="app-background" aria-hidden="true">
  <!-- decorative blobs -->
  <svg class="background-blob" style="width:900px; height:500px; left: -10%; top: -8%;" viewBox="0 0 900 500" xmlns="http://www.w3.org/2000/svg">
    <defs>
      <linearGradient id="g1" x1="0" x2="1">
        <stop offset="0%" stop-color="#3b82f6" stop-opacity="0.45" />
        <stop offset="100%" stop-color="#7c3aed" stop-opacity="0.45" />
      </linearGradient>
    </defs>
    <ellipse cx="300" cy="200" rx="420" ry="220" fill="url(#g1)" />
  </svg>

  <svg class="background-blob" style="width:700px; height:420px; right: -8%; top: 30%;" viewBox="0 0 700 420" xmlns="http://www.w3.org/2000/svg">
    <defs>
      <linearGradient id="g2" x1="0" x2="1">
        <stop offset="0%" stop-color="#7c3aed" stop-opacity="0.36" />
        <stop offset="100%" stop-color="#3b82f6" stop-opacity="0.36" />
      </linearGradient>
    </defs>
    <path d="M40,200 C120,40 400,20 620,80 C680,100 720,220 620,300 C420,370 160,320 40,260 Z" fill="url(#g2)" />
  </svg>

  <!-- very subtle palm silhouettes -->
  <svg class="background-palms" width="1400" height="400" viewBox="0 0 1400 400" xmlns="http://www.w3.org/2000/svg">
    <g fill="#24324a" transform="translate(0,40)">
      <path d="M120 160 C160 60 220 40 280 80 C340 120 360 160 320 200 C260 260 160 240 120 200 C80 160 80 140 120 160 Z" opacity="0.12" />
      <path d="M380 160 C420 60 480 40 540 80 C600 120 620 160 580 200 C520 260 420 240 380 200 C340 160 340 140 380 160 Z" opacity="0.08" />
      <path d="M980 160 C1020 60 1080 40 1140 80 C1200 120 1220 160 1180 200 C1120 260 1020 240 980 200 C940 160 940 140 980 160 Z" opacity="0.06" />
    </g>
  </svg>
</div>
