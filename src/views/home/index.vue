<template>
   <div class="container">
      <IndexBanner />

      <Bubbles />

      <Intro1 />

      <Intro2 />
      <Intro3 />
      <!-- 引导 -->
      


   </div>

</template>

<script setup>
import IndexBanner from './banner.vue';
import Intro1 from './intro_1.vue';
import Intro2 from './intro_2.vue';
import Intro3 from './intro_3.vue';
import Bubbles from '../components/bubbles.vue';
</script>



<style lang="scss" scoped>
* {
   /* 隐藏默认的滚动条样式 */
   scrollbar-width: none;
   /* Firefox */
   -ms-overflow-style: none;
   /* IE and Edge */
}


.container {
   height: 95vh;
   overflow-y: scroll;
   scroll-snap-type: y mandatory;
}

</style>