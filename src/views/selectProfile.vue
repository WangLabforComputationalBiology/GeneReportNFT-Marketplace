<template>
    <div class="card">
        <div class="card-header">{{ code }}</div>
        <div class="card-body">
            <div v-if="profiles.length > 0">
                <p>请选择一个 Profile:</p>
                <label v-for="(profile, index) in profiles" :key="index" class="radio-label">
                    <input type="radio" v-model="selectedProfile" :value="profile" />
                    {{ profile }}
                </label>
            </div>
            <div v-else>
                <p>没有可用的 Profile。</p>
            </div>
            <p>选中的 Profile: {{ selectedProfile }}</p>
            <p>{{ content }}</p>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        content: {
            type: String,
            default: '内容'
        }
    },
    data() {
        console.log('Route params in data:', this.$route.params);
        return {
            code: this.$route.params.lastSegment || '默认标题',
            profiles: [],
            selectedProfile: null
        };
    },
    watch: {
        // Watch for changes in the route parameter
        '$route.params.lastSegment': function (newSegment) {
            this.code = newSegment || '默认标题';
            this.fetchData(); // 重新获取数据
        }
    },
    mounted() {
        console.log('Mounted title:', this.code); // 打印挂载时的 title 值
        console.log('Environment variable VITE_APP_BASE_URL:', import.meta.env.VITE_APP_BASE_URL); // 打印环境变量
        this.fetchData(); // 在组件挂载时发起请求
    },
    methods: {
        async fetchData() {
            try {
                const response = await fetch(`${import.meta.env.VITE_APP_BASE_URL}/user/getProfileIds?code=${this.code}`);
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const data = await response.json();
                this.profiles = data.profiles || [];
                console.log('Fetched data:', this.profiles);
                // 你可以在这里更新 content 或其他数据
                this.content = data.content || '默认内容'; // 假设服务器返回的数据中有一个 content 字段
            } catch (error) {
                console.error('Error fetching data:', error);
                this.content = '请求失败，请重试'; // 更新 content 以显示错误信息
            }
        }
    }
}
</script>

<style scoped>
.radio-label {
    display: block;
    margin-bottom: 10px;
}
</style>