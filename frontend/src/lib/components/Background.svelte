<script lang="ts">
  import { theme } from '$lib/stores/theme'
  
  $: isDark = $theme.mode === 'dark'
  $: palmColor = isDark ? 'rgba(59, 130, 246, 0.03)' : 'rgba(10, 15, 28, 0.04)'
  $: blobColor1 = isDark ? 'rgba(59, 130, 246, 0.08)' : 'rgba(59, 130, 246, 0.12)'
  $: blobColor2 = isDark ? 'rgba(124, 58, 237, 0.06)' : 'rgba(124, 58, 237, 0.10)'
</script>

<div class="background-container" aria-hidden="true">
  <!-- Gradient Blobs -->
  <svg class="background-blob" style="width:900px; height:600px; left: -5%; top: -10%;" viewBox="0 0 900 600" xmlns="http://www.w3.org/2000/svg">
    <defs>
      <linearGradient id="blob1" x1="0" x2="1" y1="0" y2="1">
        <stop offset="0%" stop-color={blobColor1} />
        <stop offset="100%" stop-color={blobColor2} />
      </linearGradient>
    </defs>
    <ellipse cx="400" cy="280" rx="450" ry="280" fill="url(#blob1)" />
  </svg>

  <svg class="background-blob" style="width:800px; height:500px; right: -8%; top: 40%;" viewBox="0 0 800 500" xmlns="http://www.w3.org/2000/svg">
    <defs>
      <linearGradient id="blob2" x1="0" x2="1" y1="0" y2="1">
        <stop offset="0%" stop-color={blobColor2} />
        <stop offset="100%" stop-color={blobColor1} />
      </linearGradient>
    </defs>
    <path d="M50,200 C150,50 450,30 700,100 C750,130 800,250 700,350 C500,450 200,380 50,280 Z" fill="url(#blob2)" />
  </svg>

  <!-- Palm and Tropical Silhouettes -->
  <svg class="tropical-silhouettes" width="100%" height="100%" viewBox="0 0 1600 800" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid slice">
    <g fill={palmColor}>
      <!-- Palm Tree Left -->
      <path d="M150,600 Q140,550 155,520 Q165,490 150,460 Q160,420 180,400 Q200,420 190,460 Q205,490 185,520 Q195,550 180,600 Z" opacity="0.5" />
      <path d="M150,600 Q160,550 145,520 Q135,490 150,460 Q140,420 120,400 Q100,420 110,460 Q95,490 115,520 Q105,550 120,600 Z" opacity="0.4" />
      <ellipse cx="150" cy="605" rx="8" ry="5" opacity="0.6" />
      
      <!-- Palm Fronds Spreading Left -->
      <path d="M150,410 Q100,380 60,390 Q40,400 50,420 Q80,430 120,420" opacity="0.35" />
      <path d="M150,410 Q200,380 240,390 Q260,400 250,420 Q220,430 180,420" opacity="0.35" />
      
      <!-- Coconuts Left -->
      <ellipse cx="140" cy="425" rx="8" ry="10" opacity="0.25" />
      <ellipse cx="160" cy="420" rx="7" ry="9" opacity="0.25" />

      <!-- Palm Tree Right -->
      <path d="M1400,650 Q1390,600 1405,570 Q1415,540 1400,510 Q1410,470 1430,450 Q1450,470 1440,510 Q1455,540 1435,570 Q1445,600 1430,650 Z" opacity="0.4" />
      <path d="M1400,650 Q1410,600 1395,570 Q1385,540 1400,510 Q1390,470 1370,450 Q1350,470 1360,510 Q1345,540 1365,570 Q1355,600 1370,650 Z" opacity="0.3" />
      <ellipse cx="1400" cy="655" rx="8" ry="5" opacity="0.5" />

      <!-- Additional Tropical Leaves Center -->
      <path d="M800,250 Q750,220 720,240 Q700,260 720,280 Q760,290 800,270" opacity="0.2" />
      <path d="M800,250 Q850,220 880,240 Q900,260 880,280 Q840,290 800,270" opacity="0.2" />

      <!-- Small Coconut Cluster Center -->
      <ellipse cx="790" cy="265" rx="6" ry="8" opacity="0.15" />
      <ellipse cx="810" cy="268" rx="5" ry="7" opacity="0.15" />
      
      <!-- Jungle Leaf Bottom Right -->
      <path d="M1200,700 Q1150,680 1130,710 Q1120,740 1145,755 Q1180,765 1210,745" opacity="0.25" />
      
      <!-- Small Palm Bottom Center -->
      <path d="M750,720 Q745,690 755,675 Q760,665 755,655 Q760,645 770,640 Q780,645 775,655 Q780,665 775,675 Q785,690 780,720 Z" opacity="0.3" />
    </g>
  </svg>
</div>

<style lang="css">
  .background-container {
    position: fixed;
    inset: 0;
    pointer-events: none;
    z-index: -1;
    overflow: hidden;
  }

  .background-blob {
    position: absolute;
    filter: blur(80px) saturate(150%);
    opacity: 1;
    animation: float 20s ease-in-out infinite;
  }

  .background-blob:nth-child(2) {
    animation: float 25s ease-in-out infinite reverse;
  }

  .tropical-silhouettes {
    position: absolute;
    inset: 0;
    filter: blur(2px);
    opacity: 1;
  }

  @keyframes float {
    0%, 100% {
      transform: translate(0, 0) scale(1);
    }
    33% {
      transform: translate(20px, -15px) scale(1.05);
    }
    66% {
      transform: translate(-15px, 10px) scale(0.95);
    }
  }
</style>