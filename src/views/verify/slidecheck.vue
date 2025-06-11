<template>
    <div class="slider-check-box">
        <div class="slider-check" :class="rangeStatus ? 'success' : ''">
            <i @mousedown="rangeMove" :class="rangeStatus ? successIcon : startIcon"></i>
            {{ rangeStatus ? successText : startText }}
        </div>
    </div>
</template>

<script>
export default {
    props: {
        successFun: { type: Function }, // 成功回调函数
        successIcon: { type: String, default: 'el-icon-success' }, // 成功图标
        successText: { type: String, default: 'Success' }, // 成功文字
        startIcon: { type: String, default: 'el-icon-d-arrow-right' }, // 开始图标
        startText: { type: String, default: 'Silde to verify' }, // 开始文字
        errorFun: { type: Function }, // 失败回调函数
        status: { type: String } // 状态监听
    },
    data() {
        return {
            disX: 0,
            rangeStatus: false
        };
    },
    methods: {
        rangeMove(e) {
            // 如果已验证，禁止进一步操作
            if (this.rangeStatus) return;

            let ele = e.target;
            let startX = e.clientX;
            let eleWidth = ele.offsetWidth;
            let parentWidth = ele.parentElement.offsetWidth;
            let MaxX = parentWidth - eleWidth;

            // 定义滑动事件处理
            const onMouseMove = (e) => {
                let endX = e.clientX;
                this.disX = endX - startX;

                // 限制滑块移动范围
                if (this.disX <= 0) {
                    this.disX = 0;
                }
                if (this.disX >= MaxX) {
                    this.disX = MaxX;
                }

                ele.style.transition = '.1s all';
                ele.style.transform = `translateX(${this.disX}px)`;
                e.preventDefault();
            };

            // 定义鼠标释放事件处理
            const onMouseUp = () => {
                // 清理事件监听器
                document.removeEventListener('mousemove', onMouseMove);
                document.removeEventListener('mouseup', onMouseUp);

                if (this.disX < MaxX) {
                    // 验证失败，重置滑块
                    ele.style.transition = '.5s all';
                    ele.style.transform = 'translateX(0)';
                    if (this.errorFun) this.errorFun();
                } else {
                    // 验证成功
                    this.rangeStatus = true;
                    if (this.status) {
                        this.$parent[this.status] = true;
                    }
                    if (this.successFun) this.successFun();
                    // 通知父组件
                    this.$emit('ready', this.rangeStatus);

                }
            };

            // 绑定事件监听器
            document.addEventListener('mousemove', onMouseMove);
            document.addEventListener('mouseup', onMouseUp);

        }
    }
};
</script>
<style lang="scss" scoped>
$green: #169608;

@mixin jc-flex {
    display: flex;
    justify-content: center;
    align-items: center;
}

.slider-check-box {
    .slider-check {
        background-color: #e9e9e9;
        border-radius: 10px;

        position: relative;
        transition: 1s all;
        user-select: none;
        color: #585858;
        @include jc-flex;
        height: 40px;

        &.success {
            background-color: $green;
            color: #fff;

            i {
                color: $green;
            }
        }

        i {
            position: absolute;
            left: 0;
            width: 50px;
            height: 100%;
            color: $green;
            background-color: #fff;
            border: 1px solid #d8d8d8;
            border-radius: 10px;
            cursor: pointer;
            font-size: 24px;
            @include jc-flex;
        }
    }
}
</style>