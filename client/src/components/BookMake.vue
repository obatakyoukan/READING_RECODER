<template>
 <div id="make">
  <div>
  <el-form :inline="true">
   <el-form-item label="Title">
    <el-input type="text" size="large" v-model="title"></el-input>
   </el-form-item>

   <el-form-item label="Price">
    <el-input type="text" size="large" v-model="price"></el-input>
   </el-form-item>

   <el-form-item label="Writer">
    <el-input type="text" size="large" v-model="writer"></el-input>
   </el-form-item>

   <el-form-item label="Publisher">
    <el-input type="text" size="large" v-model="publisher"></el-input>
   </el-form-item>

   <el-form-item label="Published">
    <el-input type="text" size="large" v-model="published"></el-input>
   </el-form-item>

   <el-form-item label="Url">
    <el-input type="text" size="large" v-model="url"></el-input>
   </el-form-item>
 
   
   <el-form-item label="Memo" >
    <el-input type="textarea" size="large" v-model="memo"></el-input>
   </el-form-item>
   <el-form-item>
    <el-button type="primary" @click="onclick">Register</el-button>
   </el-form-item>
  
  </el-form>
   </div>

    <div v-html="memomark" />
 <div>

<!--   {{ title }}
   {{ price }} 
   {{ writer }}
   {{ publisher }}
   {{ published }}
   {{ url }}
   -->
  </div>

 </div>
</template>

<script>
// import BookInfo from '@/components/BookInfo.vue'
// import { mapActions } from 'vuex'
// import { UPDATE_CURRENT } from '@/store/mutation-types'
 
import axios from "axios";
import marked from "marked";

 export default {
  name: 'book-make',
  components: {
 //  BookInfo
  },
  data() {
   return {
    title: '',
    price: '',
    writer: '',
    publisher: '',
    published: '',
    url: '',
    imagetype: 'pdf',
    memo: ''
   }
  },
  computed: {
  memomark: function(){
   return marked(this.memo)
   //return marked(this.form.memo)
  }
  },
  methods: {
   //...mapActions( [ UPDATE_CURRENT] ),
   onclick() {
    /*let book = {
     id: this.url,
     title: this.title,
     author: this.author,
     price : this.price ? this.price : '-',
     publisher: this.publisher,
     published: this.published,
     image: this.url,
     imagetype: this.imagetype,
     url: this.url,
     memo: this.memo
    }*/
      const params = {
    //   uid: "aaa",
       title: this.title,
       author: this.author,
       price: String(this.price),
       publisher: this.publisher,
       published: this.published,
       image: this.url,
       imagetype: this.imagetype,
       url: this.url,
       memo: this.memo
      }
      axios.post('http://localhost:8888/create_noid', params )
    //axios.post('http://localhost:8888/create_noid?title='+this.title+'&author='+this.author+'&price='+this.price+'&publisher='+this.publisher+'&published='+this.published+'&imagetype='+this.imagetype+'&url='+this.url+'&memo='+this.memo +'&image='+ this.url )
    //this[ UPDATE_CURRENT](book)
    this.$router.push('/')
   }
  }
 }

</script>


<style scoped>
 #make form {
  margin-top: 15px;
 }
</style>
