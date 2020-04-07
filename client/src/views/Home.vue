<template>
  <div class="list">
   <div><p> There are {{ bookCount }} books.</p> </div>
 <!--  <div> {{ books }} </div> -->
   <BookInfo v-for="(b, i) of books"
    :linkable="true" :book="b" :index="i+1" :key="b.isbn">  </BookInfo>
    <div>
     <el-pagination background layout="prev, pager, next"
      :total="bookCount" :page-size="3" @current-change="onchange"></el-pagination>
    </div> 
  </div> 
</template>

<script>
import BookInfo from '@/components/BookInfo.vue'
import axios from "axios";

export default {
  name: 'Home',
  data() {
   return {
    books: [],
    allbooks: [],
    init_num: 0
   }
  },
  components: {
   BookInfo
  },
  computed: {
   bookCount() {
    return this.allbooks.length
   },
  /* getRangeByPage: function(page){
    const SIZE = 3 
    return this.allbooks.slice( (page-1)*SIZE , (page-1)*SIZE+SIZE)
   }*/
  },
  methods: {
  onchange(page) {
    const SIZE = 3 
    this.books = this.allbooks.slice( (page-1)*SIZE , page*SIZE )
    //this.books = this.getRangeByPage(page)
   }
  },
  mounted() {
     axios.get("http://localhost:8888/list")
    .then( res => { return   res.data.items } )
    .then( data => {
     let allbook = []
     for(let b of data ){
       allbook.push(
       {
        id: b.uid,
	title: b.title,
	author: b.author,
	price: b.price,
	publisher: b.publisher,
	published: b.published,
	image: b.image,
	imagetype: b.imagetype,
	url: b.url,
	memo: b.memo
       }
       )
     }
     this.allbooks = allbook
     this.init_num = this.allbooks.length
     const SIZE = 3
     this.books = this.allbooks.slice( 0, SIZE )
    })
  },
  created:  function(){
   let that = this;//setInterval内ではthisがvueインスタンスからグローバルオブジェクトに変化する
   var check_time = [1000,2000,3000,5000,7000,9000,10000,12000,15000,20000];
   for(let t of check_time ){
    this.timer = setTimeout(function(){
   axios.get("http://localhost:8888/list")
    .then( res => { return   res.data.num } )
    /*.then( data => {
     let allbook = []
     for(let b of data ){
       allbook.push(
       {
        id: b.uid,
	title: b.title,
	author: b.author,
	price: b.price,
	publisher: b.publisher,
	published: b.published,
	image: b.image,
	imagetype: b.imagetype,
	url: b.url,
	memo: b.memo
       })
     }
     that.allbooks = allbook
    })*/
   .then( num => {
    if( num !== that.init_num ){
   // alert(that.init_num)
   // alert(num)
     location.reload()
    }
   })

       },t)
       }
  }
  
 
}
</script>
