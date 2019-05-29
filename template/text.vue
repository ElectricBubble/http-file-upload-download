<div id="mushroom-text">
    <el-input
            type="textarea"
            :rows="32"
            @change="handleTextChange"
            placeholder="Please input"
            v-model="mrText">
    </el-input>
</div>

<script>
    "use strict";
    let mrText = '<%==s sText %>';
    let textExtData = {
        data() {
            return {
                mrText,
            }
        },
        methods: {
            handleTextChange() {
                // console.log("文本发生了改变"+"["+this.mrText+"]");
                let request = new XMLHttpRequest();
                request.open("POST", location.href);
                request.send(this.mrText);
            },
        },
    };
    let textCtor = Vue.extend(textExtData);
    new textCtor().$mount('#mushroom-text');
    document.getElementsByTagName("title")[0].innerText = "⌨️";
</script>