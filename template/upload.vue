<div id="mushroom-upload">
    <el-upload
            class="upload-demo"
            ref="upload"
            drag
            multiple
            action="/upload/"
            :auto-upload="true">
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击选择文件</em></div>
    </el-upload>
</div>

<script>
    "use strict";
    let uploadExtData = {
        data() {
            return {};
        },
        methods: {
            submitUpload() {
                this.$refs.upload.submit();
            },
        }
    };
    let uploadPageCtor = Vue.extend(uploadExtData);
    new uploadPageCtor().$mount('#mushroom-upload');
    document.getElementsByTagName("title")[0].innerText = "🚀";
</script>

<style scoped>
    .el-upload {
        width: 100%;
        height: 512px;
    }

    .el-upload-dragger {
        width: 100%;
        height: 100%;
        border: 2px dashed #d9d9d9;
    }

    .el-upload-dragger .el-icon-upload {
        font-size: 256px;
        line-height: 100%;
    }

    .el-upload-dragger .el-upload__text {
        font-size: 32px;
    }
</style>
