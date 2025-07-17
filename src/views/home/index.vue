<template>
   <div class="container">
      <Bubbles />

      <swiper :direction="'vertical'" @swiper="onVerticalSwiperInit" :slides-per-view="1" :touchRatio="0"
         :mousewheel="true" :speed="700" @slideChange="onVerticalSlideChange" :modules="modules"
         class="fullpage-swiper">
         <div class="guide" v-if="page > 0">
            <div class="point" v-for="i in 5" :key="i" :class="{ active: page === i - 1 }" @click="toPage(i - 1)" />
         </div>
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
import { ref, watch } from 'vue'
import { Swiper, SwiperSlide } from 'swiper/vue'
import { Mousewheel, Pagination, Navigation } from 'swiper'
import IndexBanner from './banner.vue';
import Intro1 from './intro_1.vue';
import Intro2 from './intro_2.vue';
import Intro3 from './intro_3.vue';
import Intro4 from './intro_4.vue';
import Bubbles from '../components/bubbles.vue';

/* swiper默认配置项 */
const modules = [Pagination, Navigation, Mousewheel];

/* 首屏禁止上滚 */
// const allowPrev = ref(true)

/* 垂直滚动插入水平滚动 */
// let allowNext = ref(true);

/* 初始化垂直 Swiper */
const verticalSwiperPage = ref(null);
const onVerticalSwiperInit = (swiper) => {
   verticalSwiperPage.value = swiper;//获取swiper实例
   // console.log('垂直swiper初始化', verticalSwiperPage.value.activeIndex);
};
const page = ref(0);
const onVerticalSlideChange = (swiper) => {
   page.value = swiper.activeIndex;
   // console.log('垂直swiper切换', page.value);
}

const toPage = (index) => {
   verticalSwiperPage.value.slideTo(index);
}
</script>


<style lang="scss" scoped>
.container {
   height: 95vh;
   position: relative;
}

/**页数指示点 */
.guide {
   position: absolute;
   right: 10px;
   top: 50%;
   display: flex;
   gap: 15px;
   flex-direction: column;
   transform: translateY(-70%);
   animation: fadeIn 0.4s ease-in 0s forwards;

   @keyframes fadeIn {
      0% {
         opacity: 0;
         transform: translateX(20%);
      }

      100% {
         opacity: 1;
      }
   }

   .point {
      width: 12px;
      height: 12px;
      background-color: #eee;
      transform: rotate(45deg);
      border-radius: 25%;
      cursor: pointer;
      box-shadow: solid 0 0 5px #ccc;
   }

   .active {
      background-color: #169608;
   }
}

.fullpage-swiper {
   width: 100%;
   height: 100%;
}

.swiper-slide {
   width: 100%;
   height: 100%;

   &:nth-child(3) {
      //防止第三页绿色背景露出
      border-bottom: #fff solid 1px;
      border-top: #fff solid 1px;
   }
}

.horizontal-swiper {
   height: 95vh;
}
</style>