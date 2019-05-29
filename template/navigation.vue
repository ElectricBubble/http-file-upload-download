<div id="mushroom-nav">
    <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
        <el-menu-item index="1"><i class="el-icon-document"></i></el-menu-item>
        <el-menu-item index="2"><i class="el-icon-document-add"></i></el-menu-item>
        <el-menu-item index="3"><i class="el-icon-edit-outline"></i></el-menu-item>
        <el-menu-item index="4" disabled><i class="el-icon-picture-outline"></i></el-menu-item>
    </el-menu>
</div>

<script>
    "use strict";
    let urlMap = {};
    urlMap['1'] = '/mushroom/';
    urlMap['2'] = '/up';
    urlMap['3'] = '/text';

    let extData = {
        data() {
            return {
                activeIndex: '<%==i activeIndex %>',
            };
        },
        methods: {
            handleSelect(key, keyPath) {
                console.log(key, keyPath);
                if (urlMap.hasOwnProperty(key)) {
                    console.log(urlMap[key]);
                    window.location.href = urlMap[key];
                }
            }
        }
    };
    let navMenuCtor = Vue.extend(extData);
    new navMenuCtor().$mount('#mushroom-nav');
</script>