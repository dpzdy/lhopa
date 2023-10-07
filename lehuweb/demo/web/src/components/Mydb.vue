<template>
  <div id="db">
    <div class="p1">
      <div class="title">
        <span>id</span>
        <span>城市名</span>
        <span>标题</span>
        <span>价格</span>
        <span>行程</span>
        <span>年份</span>
      </div>
      <ul>
        <li v-for="item in list1" :key="item.dept_no">
          <span>{{ item.id }}</span>
          <span>{{ item.cityname}}</span>
          <span>{{ item.title }}</span>
          <span>{{ item.price }}</span>
          <span>{{ item.kilometer }}</span>
          <span>{{ item.year }}</span>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import { request } from "../utils/requset.js";
export default {
  name: "Mydb",
  data() {
    return {
      //存储get到的数据默认为空数组
      list1: [],
    };
  },
  methods: {
    async axiosPost() {
      const { data: res } = await this.$http({
        url: "/api",
        method: "get",
      });
      console.log(res);
      this.list1 = res.data
    },
  },
  created() {
    this.axiosPost();
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
span {
  width: 250px;
  height: 40px;
  display: inline-block;
}
.title {
  font-size: larger;
}
</style>
