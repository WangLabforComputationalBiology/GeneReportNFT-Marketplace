<template>
    <body>
        <div style="max-width: 600px">
            <el-alert
            title="Success alert"
            type="success"
            description="More text description"
            show-icon
            />
        </div>
        <h1>{{ code }}</h1>
        <div class="banner">
            <span class="banner-title">Create</span><span class="tip">Please select a profile:</span>
        </div>

        <div class="card-body">
            <div v-if="profiles.length > 0">
                <p>请选择一个 Profile:</p>
                <p>---------</p>
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
        <button @click="authorizeProfile">授权报告</button>
    </body>
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
            selectedProfile: null,
            showAlert: false // 添加状态变量
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
                const response = await fetch(`http://127.0.0.1:4523/m1/4576706-4225408-default/user/getProfileIds`);
                //const response = await fetch(`${import.meta.env.VITE_APP_BASE_URL}/user/getProfileIds?code=${this.code}`);
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
        },
        async authorizeProfile() {
            // 在这里处理授权逻辑
            try {
            const response = await fetch(`${import.meta.env.VITE_APP_BASE_URL}/user/saveProfile`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    code: this.code,
                    profileId: this.selectedProfile
                })
            });

            if (response.ok) { // 检查响应状态码是否为 200
                this.showAlert = true; // 显示 alert
                console.log('Profile authorized successfully');
            } else {
                console.error('Failed to authorize profile:', response.statusText);
            }
        } catch (error) {
            console.error('Error authorizing profile:', error);
        }

           
        }
    }
}
</script>



<style lang="scss" scoped>

.el-alert {
  margin: 20px 0 0;
}
.el-alert:first-child {
  margin: 0;
}
body {
    margin: auto;
    width: 1400px;
    min-height: calc(100vh - 80px);
    width: 80vw;
    min-width: 1200px;
    overflow: visible;
}

.banner {
    .banner-title {
        font-size: 70px;
        color: #67C23A;
    }

    .tip {
        margin-left: 20px;
        font-size: 20px;
        color: #99a9bf;
    }
}

.radio-label {
    display: block; // 每个选项独立成行
    margin-bottom: 33px;
    cursor: pointer; // 鼠标悬停时显示指针

    &:hover {
        font-size: 1.7em; // 字体放大
        color: blue; // 字体颜色变蓝
    }

    input[type="radio"] {
        margin-right: 10px; // 单选按钮和文本之间的间距
    }
}
</style>