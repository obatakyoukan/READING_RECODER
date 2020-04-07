<template>
 <div id="search">
  <el-form :inline="true">
   <el-form-item label="Keyword">
    <el-input type="text" size="large" v-model="keyword"></el-input>
   </el-form-item>

   <el-form-item label="Title">
    <el-input type="text" size="large" v-model="title_book"></el-input>
   </el-form-item>

   <el-form-item label="Author">
    <el-input type="text" size="large" v-model="author_book"></el-input>
   </el-form-item>

   <el-form-item>
    <el-button type="primary" @click="onclick">Search</el-button>
   </el-form-item>
    検索ヒット件数:{{num_book}}
  </el-form>
  <hr />
  <BookInfo v-for="(b, i) of books"
   :linkable="true" :book="b" :index="i+1" :key="b.isbn"></BookInfo>
 </div>
</template>

<script>
 import BookInfo from '@/components/BookInfo.vue'
 
 export default {
  name: 'book-search',
  components: {
   BookInfo
  },
  data() {
   return {
    keyword: '',
    title_book: '',
    author_book: '',
    num_book: 0,
    books: []
   }
  },
  methods: {
   onclick: function() {
    let keywords = this.keyword.replace(' ','+').replace('　','+').replace(',','+')
    let title_books = this.title_book ? '+intitle:'+this.title_book : ''
    let author_books = this.author_book ? '+inauthor:'+this.author_book : ''
    let bookurl = 'https://www.googleapis.com/books/v1/volumes?q=' +  keywords + title_books + author_books
    this.$http( bookurl )
      .then( (response) => {
       return response.json()
      })
      .then( (data) => {
       this.books = []
       this.num_book = data.totalItems
       for( let b of data.items){
        let authors = b.volumeInfo.authors
	let price = b.saleInfo.listPrice
	let img = b.volumeInfo.imageLinks
	this.books.push({
	id: b.id,
	title: b.volumeInfo.title,
	author: authors ? authors.join(',') : '',
	price: price ? price.amount : '-',
	publisher: b.volumeInfo.publisher,
	published: b.volumeInfo.publishedDate,
	image: img ? img.smallThumbnail : '',
	imagetype: 'png',
	url: b.volumeInfo.infoLink,
	memo: ""
       })
       }
      })
   }
  }
 }
</script>

<style scoped>
 #search form {
  margin-top: 15px;
 }
</style>
