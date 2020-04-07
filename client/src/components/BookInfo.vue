<template>
<div class="clearfix" :class="{ linkable }" @click="onclick">
 <div v-if="book.imagetype==='pdf'"> <div class="image"><iframe :src="book.image" width="200" height="200" /> </div> </div>
 <div v-else><div class="image"><img :src="book.image" /> </div> </div>
 <div class="details">
  <ul>
   <li v-if="index">{{ index }}. </li>
   <li> Title:     {{book.title}} </li>
   <li> Price:     {{book.price}} </li>
   <li> Writer:    {{book.author}}</li>
   <li> Publisher: {{book.publisher}} </li>
   <li> Published: {{book.published}} </li>
   <li> imagetype: {{book.imagetype}} </li>
   <li> <a :href="book.url">Link</a>  </li>
 </ul>
 </div>
</div>
</template>

<script>
//bookを決定してformに飛ぶためのメソッドのみなので直さない
import { mapActions } from 'vuex'
import { UPDATE_CURRENT } from '@/store/mutation-types'

export default {
 name: 'book-info',
 props: {
  index: { type: Number },
  linkable: { type: Boolean, default: false },
  book: { type: Object }
 },
 methods: {
  ...mapActions([UPDATE_CURRENT]),
  onclick() {
   if (this.linkable) {
    this[UPDATE_CURRENT](this.book)
    this.$router.push('/form')
   }
  }
 }
}
</script>

<style scoped>
.linkable:hover {
    cursor: pointer;
    background-color: #ff9;
}

.image {
  float:left;
}

.image img {
  height: 180px;
}

.details {
  float: left;
}

.details li {
  list-style-type: none;
  text-align: left;
}

.clearfix {
  width: 100%;
  margin-bottom: 10px;
}

.clearfix:after {
  content: "";
  display: block;
  clear: both;
}
</style>
