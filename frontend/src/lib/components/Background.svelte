<script lang="ts">
  import { theme } from '$lib/stores/theme'
  import { onMount } from 'svelte'
  import { browser } from '$app/environment'
  
  $: isDark = $theme.mode === 'dark'
  $: palmColor = isDark ? 'rgba(59, 130, 246, 0.03)' : 'rgba(10, 15, 28, 0.04)'
  $: blobColor1 = isDark ? 'rgba(59, 130, 246, 0.08)' : 'rgba(59, 130, 246, 0.12)'
  $: blobColor2 = isDark ? 'rgba(124, 58, 237, 0.06)' : 'rgba(124, 58, 237, 0.10)'
  
  let flowerColor = '#ec4899'

  // Trigger page transition animation
  onMount(() => {
    if (browser) {
      const bg = document.querySelector('.tropical-background')
      if (bg) {
        bg.classList.add('page-transition')
        setTimeout(() => bg?.classList.remove('page-transition'), 500)
      }
    }
  })
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

  <!-- Large Palm Tree Left -->
  <svg class="palm palm-left palm-sway" viewBox="0 0 200 400" xmlns="http://www.w3.org/2000/svg">
    <g fill={palmColor}>
      <!-- Trunk -->
      <path d="M95,400 Q90,350 95,300 Q100,250 95,200 Q90,150 95,100 Q100,50 100,0" stroke={palmColor} stroke-width="8" fill="none" opacity="0.5"/>
      
      <!-- Left fronds -->
      <path d="M100,50 Q60,30 20,40 Q10,45 15,55 Q40,60 70,55" opacity="0.4"/>
      <path d="M100,80 Q50,60 10,80 Q5,90 12,98 Q50,100 85,90" opacity="0.35"/>
      <path d="M100,110 Q55,95 15,120 Q8,130 18,140 Q60,135 90,125" opacity="0.3"/>
      
      <!-- Right fronds -->
      <path d="M100,50 Q140,30 180,40 Q190,45 185,55 Q160,60 130,55" opacity="0.4"/>
      <path d="M100,80 Q150,60 190,80 Q195,90 188,98 Q150,100 115,90" opacity="0.35"/>
      <path d="M100,110 Q145,95 185,120 Q192,130 182,140 Q140,135 110,125" opacity="0.3"/>
      
      <!-- Coconuts -->
      <ellipse cx="85" cy="100" rx="6" ry="8" opacity="0.3"/>
      <ellipse cx="95" cy="105" rx="5" ry="7" opacity="0.25"/>
      <ellipse cx="105" cy="102" rx="6" ry="8" opacity="0.3"/>
    </g>
  </svg>

  <!-- Palm Tree Right -->
  <svg class="palm palm-right palm-sway" style="animation-delay: -2s;" viewBox="0 0 200 400" xmlns="http://www.w3.org/2000/svg">
    <g fill={palmColor}>
      <!-- Trunk -->
      <path d="M105,400 Q110,350 105,300 Q100,250 105,200 Q110,150 105,100 Q100,50 100,0" stroke={palmColor} stroke-width="8" fill="none" opacity="0.4"/>
      
      <!-- Fronds -->
      <path d="M100,60 Q70,45 30,55 Q22,62 28,70 Q55,72 85,65" opacity="0.35"/>
      <path d="M100,60 Q130,45 170,55 Q178,62 172,70 Q145,72 115,65" opacity="0.35"/>
      <path d="M100,95 Q65,82 25,100 Q18,110 26,118 Q65,115 92,105" opacity="0.3"/>
      <path d="M100,95 Q135,82 175,100 Q182,110 174,118 Q135,115 108,105" opacity="0.3"/>
      
      <!-- Coconuts -->
      <ellipse cx="90" cy="95" rx="5" ry="7" opacity="0.25"/>
      <ellipse cx="110" cy="98" rx="5" ry="7" opacity="0.25"/>
    </g>
  </svg>

  <!-- Flowers Bottom Left -->
  <svg class="flowers flowers-left flower-float" viewBox="0 0 300 200" xmlns="http://www.w3.org/2000/svg">
    <g fill={flowerColor}>
      <!-- Hibiscus-style flower -->
      <path d="M50,100 Q30,80 40,60 Q60,50 80,60 Q90,40 110,50 Q120,70 110,90 Q130,100 120,120 Q100,130 80,120 Q70,140 50,130 Q40,110 50,100" opacity="0.25"/>
      <circle cx="80" cy="95" r="8" opacity="0.2"/>
      
      <!-- Tropical leaves -->
      <path d="M150,120 Q130,100 140,80 Q160,75 175,90 Q180,110 165,125 Q145,130 150,120" opacity="0.2"/>
      <path d="M170,140 Q155,125 160,110 Q175,105 185,118 Q188,133 178,143 Q165,146 170,140" opacity="0.18"/>
      
      <!-- Small flowers -->
      <circle cx="220" cy="110" r="12" opacity="0.15"/>
      <circle cx="240" cy="130" r="10" opacity="0.12"/>
      <circle cx="200" cy="150" r="8" opacity="0.1"/>
    </g>
  </svg>

  <!-- Flowers Bottom Right -->
  <svg class="flowers flowers-right flower-float" style="animation-delay: -3s;" viewBox="0 0 300 200" xmlns="http://www.w3.org/2000/svg">
    <g fill={flowerColor}>
      <!-- Large tropical flower -->
      <path d="M150,100 Q170,80 190,90 Q200,110 185,125 Q165,135 145,125 Q135,105 150,100" opacity="0.22"/>
      <circle cx="167" cy="108" r="10" opacity="0.18"/>
      
      <!-- Monstera leaf -->
      <path d="M80,120 Q60,100 70,80 Q90,70 110,85 Q120,105 105,120 Q85,128 80,120" opacity="0.2"/>
      
      <!-- Bird of Paradise shape -->
      <path d="M230,140 Q210,125 215,105 Q235,95 250,110 Q255,130 240,145 Q225,150 230,140" opacity="0.18"/>
      
      <!-- Small decorative flowers -->
      <circle cx="50" cy="150" r="9" opacity="0.12"/>
      <circle cx="30" cy="130" r="7" opacity="0.1"/>
    </g>
  </svg>

  <!-- Jungle Vines Top -->
  <svg class="vines vines-top" viewBox="0 0 1200 300" xmlns="http://www.w3.org/2000/svg">
    <g fill="none" stroke={palmColor} stroke-width="2" opacity="0.15">
      <path d="M0,50 Q100,30 200,50 T400,50 T600,50" />
      <path d="M200,20 Q300,10 400,25 T600,25 T800,25" opacity="0.12"/>
      <path d="M400,80 Q500,60 600,75 T800,75 T1000,75" opacity="0.1"/>
      
      <!-- Hanging leaves -->
      <ellipse cx="150" cy="60" rx="8" ry="15" fill={flowerColor} opacity="0.15"/>
      <ellipse cx="350" cy="35" rx="6" ry="12" fill={flowerColor} opacity="0.12"/>
      <ellipse cx="550" cy="85" rx="7" ry="14" fill={flowerColor} opacity="0.1"/>
      <ellipse cx="750" cy="55" rx="8" ry="16" fill={flowerColor} opacity="0.13"/>
    </g>
  </svg>

  <!-- Additional decorative elements -->
  <svg class="decorative decorative-center flower-float" style="animation-delay: -4.5s;" viewBox="0 0 400 300" xmlns="http://www.w3.org/2000/svg">
    <g fill={flowerColor}>
      <!-- Fern fronds -->
      <path d="M200,150 Q180,130 190,110 Q210,105 220,120 Q215,140 200,150" opacity="0.15"/>
      <path d="M200,150 Q220,130 210,110 Q190,105 180,120 Q185,140 200,150" opacity="0.15"/>
      
      <!-- Tropical berries -->
      <circle cx="150" cy="100" r="5" opacity="0.12"/>
      <circle cx="160" cy="95" r="4" opacity="0.1"/>
      <circle cx="155" cy="108" r="4" opacity="0.1"/>
      
      <circle cx="280" cy="180" r="6" opacity="0.14"/>
      <circle cx="290" cy="175" r="5" opacity="0.12"/>
      <circle cx="285" cy="190" r="4" opacity="0.1"/>
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
    background: var(--bg);
    transition: background-color 0.3s ease;
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

  @keyframes palm-sway {
    0%, 100% {
      transform: rotate(0deg);
    }
    25% {
      transform: rotate(1deg);
    }
    75% {
      transform: rotate(-1deg);
    }
  }

  @keyframes flower-float {
    0%, 100% {
      transform: translateY(0px);
    }
    50% {
      transform: translateY(-20px);
    }
  }

  /* Palm trees */
  .palm {
    position: absolute;
    opacity: 1;
    filter: blur(0.5px);
  }

  .palm-left {
    left: 5%;
    bottom: 0;
    height: 60vh;
    max-height: 500px;
    width: auto;
    transform-origin: bottom center;
  }

  .palm-right {
    right: 8%;
    bottom: 0;
    height: 50vh;
    max-height: 420px;
    width: auto;
    transform-origin: bottom center;
  }

  .palm-sway {
    animation: palm-sway 4s ease-in-out infinite;
  }

  /* Flowers */
  .flowers {
    position: absolute;
    opacity: 1;
    filter: blur(0.3px);
  }

  .flowers-left {
    left: 15%;
    bottom: 5%;
    height: 180px;
    width: auto;
  }

  .flowers-right {
    right: 18%;
    bottom: 8%;
    height: 160px;
    width: auto;
  }

  .flower-float {
    animation: flower-float 6s ease-in-out infinite;
  }

  /* Vines */
  .vines-top {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 200px;
    width: 100%;
    opacity: 1;
  }

  /* Decorative center elements */
  .decorative-center {
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    height: 250px;
    width: auto;
    opacity: 0.8;
  }

  /* Responsive */
  @media (max-width: 1024px) {
    .palm-left {
      left: 2%;
      height: 50vh;
    }

    .palm-right {
      right: 3%;
      height: 40vh;
    }

    .flowers-left {
      left: 8%;
      height: 140px;
    }

    .flowers-right {
      right: 10%;
      height: 120px;
    }
  }

  @media (max-width: 640px) {
    .palm-left,
    .palm-right {
      height: 35vh;
    }

    .flowers-left,
    .flowers-right {
      height: 100px;
    }

    .decorative-center {
      height: 180px;
      opacity: 0.6;
    }

    .vines-top {
      height: 120px;
    }
  }

  /* Hide on very small screens to improve performance */
  @media (max-width: 375px) {
    .decorative-center,
    .vines-top {
      display: none;
    }
  }
</style>
