<template>
  <div id="db">
    <div class="p1">
      <div class="title">
        <span>姓名</span>
        <span>地址</span>
        <span>学院</span>
      </div>
      <ui>
        <li v-for="item in list1" :key="item.dept_no">
          <!-- <span>{{ item.avatar }}</span> -->
          <img :src="item.imgURL" alt="">
          <span>{{ item.hot }}</span>
          <span>{{ item.cmtNum }}</span>
        </li>
      </ui>
    </div>

  </div>
</template>

<script>
import {request} from '../utils/requset.js'
export default {
  name: "Mydb",
  data() {
    return {
      //存储get到的数据默认为空数组
      list1: [],
    };
  },
  methods: {
    async initData() {
      // 请求回来的数据在页面要用到，必须转存到data中
      // const { data: res1 } = await this.$http.get("/api/dept");
      const { data: res1 } = await request.get("/api");
      console.log(res1);
      this.list1 = res1.data;
    },

  },
  created() {
    this.initData();
  },
};
</script>

<style lang="less">
#db {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;

  margin-top: 60px;
}
li {
  list-style: none;
}
img{
  width: 10%;
  height: 10%;
  border-radius: 50%;
  z-index: 999;
}
span {
  width: 30%;
  height: 40px;
  display: inline-block;
}
.title {
  font-size: larger;
}
</style>
