<script lang="ts">
  import { theme } from '$lib/stores/theme'
  import { onMount } from 'svelte'
  import { browser } from '$app/environment'
  
  $: isDark = $theme.mode === 'dark'
  $: palmColor = isDark ? 'rgba(255, 255, 255, 0.04)' : 'rgba(0, 0, 0, 0.03)'
  $: flowerColor = isDark ? 'rgba(255, 255, 255, 0.03)' : 'rgba(0, 0, 0, 0.02)'
  
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

<div class="tropical-background" aria-hidden="true">
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

<style>
  .tropical-background {
    position: fixed;
    inset: 0;
    pointer-events: none;
    z-index: 0;
    overflow: hidden;
    background: var(--bg);
    transition: background-color 0.3s ease;
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