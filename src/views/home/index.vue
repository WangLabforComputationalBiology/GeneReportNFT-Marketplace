<template>
   <div class="container">
      <Bubbles />
      <swiper :direction="'vertical'" :slides-per-view="1" :mousewheel="true" :speed="800" @slideChange="onSlideChange" :allowSlidePrev="allowPrev"
         :modules="modules" class="fullpage-swiper">
         <Swiper-slide>
            <IndexBanner />
         </Swiper-slide>

         <Swiper-slide>
            <Intro1 />
         </Swiper-slide>

         <Swiper-slide>
            <Intro2 />
         </Swiper-slide>

         <Swiper-slide>
            <Intro3 />
         </Swiper-slide>

         <Swiper-slide>
            <Intro4 />
         </Swiper-slide>
      </swiper>
   </div>

</template>

<script setup>
import { ref } from 'vue'
import { Swiper, SwiperSlide } from 'swiper/vue'
import { Mousewheel, Pagination, Navigation } from 'swiper'
import 'swiper/css'
import 'swiper/css/pagination'
import 'swiper/css/navigation'
import IndexBanner from './banner.vue';
import Intro1 from './intro_1.vue';
import Intro2 from './intro_2.vue';
import Intro3 from './intro_3.vue';
import Intro4 from './intro_4.vue';
import Bubbles from '../components/bubbles.vue';

const modules = [Pagination, Navigation, Mousewheel];

/* 首屏禁止上滚 */
const allowPrev = ref(false)
const onSlideChange= (swiper) => {
   allowPrev.value = swiper.activeIndex > 0
}
</script>



<style lang="scss" scoped>
.container {
   height: 95vh;
}

.fullpage-swiper {
   width: 100%;
   height: 100%;
}

.swiper-slide {
   width: 100%;
   height: 100%;
}
</style>