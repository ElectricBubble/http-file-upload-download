<div id="mushroom-fs">
    <el-table
            :data="fsItems"
            stripe
            empty-text="此文件夹无文件"
            style="width: 100%">
        <el-table-column
                prop="name">
        </el-table-column>
        <el-table-column
                width="76"
                prop="size">
        </el-table-column>
        <el-table-column
                width="124">
            <template slot-scope="scope">
                <el-button icon="el-icon-folder-opened" @click="handleClick(scope.row)" :disabled="!scope.row.isDir" circle></el-button>
                <el-button icon="el-icon-download" @click="handleClick(scope.row)" :disabled="scope.row.isDir" circle></el-button>
            </template>
        </el-table-column>
    </el-table>
</div>

<script>
    "use strict";
    let fsExtData = {
        methods: {
            handleClick(row) {
                if (row.isDir) {
                    let dirURL = location.href;
                    if (location.href.endsWith("/")) {
                        dirURL += row.name;
                    } else {
                        dirURL += "/" + row.name;
                    }
                    location.assign(dirURL);
                } else {
                    let dlURL = location.origin + "/download/" + row.filename;
                    let request = new XMLHttpRequest();
                    request.open("GET", dlURL);
                    request.send();
                    location.assign(dlURL);
                }
            }
        },
        data() {
            return {
                fsItems: [
                    <% for _, v := range fsItems { %>
                    {
                        name: '<%==s v.Name %>',
                        isDir: <%==b v.IsDir %>,
                        filename: '<%==s v.Filename %>',
                        size: '<%==s v.Size %>',
                    },
                    <% } %>
                ]
            }
        }
    };
    let fsCtor = Vue.extend(fsExtData);
    new fsCtor().$mount('#mushroom-fs');
    document.getElementsByTagName("title")[0].innerText = "📄";
</script>