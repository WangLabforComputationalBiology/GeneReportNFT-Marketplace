<template>
   <div class="container">
      <Bubbles />
      <swiper :direction="'vertical'" @swiper="onVerticalSwiperInit" :slides-per-view="1" :touchRatio="0"
         :mousewheel="true" :speed="800" @slideChange="onVerticalSlideChange" :allowSlidePrev="allowPrev"
         :allowSlideNext="allowNext" :modules="modules" class="fullpage-swiper">
         <Swiper-slide>
            <IndexBanner />
         </Swiper-slide>

         <Swiper-slide>
            <swiper ref="horizontalSwiper" @swiper="onHorizontalSwiperInit" :modules="modules" :direction="'horizontal'"
               :slides-per-view="1" :free-mode="true" @slideChange="onHorizontalSlideChange" :mousewheel="true"
               :allowSlidePrev="allowPrev2" :speed="800" class="horizontal-swiper">

               <Swiper-slide>
                  <Intro1 />
               </Swiper-slide>

               <Swiper-slide>
                  <Intro1_2 />
               </Swiper-slide>
            </swiper>
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
import { ref, } from 'vue'
import { Swiper, SwiperSlide } from 'swiper/vue'
import { Mousewheel, Pagination, Navigation } from 'swiper'
import 'swiper/css'
import 'swiper/css/pagination'
import 'swiper/css/navigation'
import IndexBanner from './banner.vue';
import Intro1 from './intro_1.vue';
import Intro1_2 from './intro_1_2.vue';
import Intro2 from './intro_2.vue';
import Intro3 from './intro_3.vue';
import Intro4 from './intro_4.vue';
import Bubbles from '../components/bubbles.vue';

/* swiper默认配置项 */
const modules = [Pagination, Navigation, Mousewheel];

/* 首屏禁止上滚 */
const allowPrev = ref(true)

/* 垂直滚动插入水平滚动 */
let allowNext = ref(true);
let verticalSwiper = ref(null);
let horizontalSwiper = ref(null);

/* 初始化垂直 Swiper */
const onVerticalSwiperInit = (swiper) => {
   verticalSwiper.value = swiper;
};

/* 初始化水平 Swiper */
const onHorizontalSwiperInit = (swiper) => {
   horizontalSwiper.value = swiper;
};

/* 水平滑动变化时 */
const onHorizontalSlideChange = (swiper) => {
   // 监听滚轮事件获取滚轮状态
   swiper.el.onwheel = (e) => {
      if (e.deltaY > 0 && swiper.isEnd) {
         setTimeout(() => {
            // e.deltaY < 0 表示向上滚动，e.deltaY > 0 表示向下滚动
            allowNext.value = true;
         }, 500)
      }
      if (e.deltaY < 0 && swiper.isBeginning) {
         allowPrev.value = true;
      }
   };

};

/* 垂直滑动变化时 */
const onVerticalSlideChange = (swiper) => {
   allowPrev.value = verticalSwiper.value.isBeginning;
   allowNext.value = verticalSwiper.value.isBeginning;

   if (swiper.activeIndex == 1) { // 假设垂直滚动到第2页（索引1）
      allowPrev.value = true;// 启用垂直滚动
      allowNext.value = false;// 禁用垂直滚动
   } else {
      allowPrev.value = true;// 启用垂直滚动
      allowNext.value = true;// 启用垂直滚动
   }
};
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

   &:nth-child(3) {
      border: #fff solid 1px;
   }
}

.horizontal-swiper {
   height: 95vh;
}
</style>