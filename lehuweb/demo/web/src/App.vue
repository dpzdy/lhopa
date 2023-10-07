<template>
  <div id="app">
    <Mydb></Mydb>
   <!-- <Getdata></Getdata> -->

  </div>
</template>

<script>
import Mydb from './components/Mydb.vue'
import Getdata from './components/Getdata.vue'
export default {
  name: "App",
  components: {
    Mydb,
    Getdata
  },
  
};
</script>

