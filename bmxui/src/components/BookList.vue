<template>
  <div>
    <h2>Book List</h2>
    <ul>
      <li v-for="book in books" :key="book.id">
        <span>{{ book.title }} by {{ book.author }}</span>
        <button @click="editBook(book)">Edit</button>
        <button @click="deleteBook(book.id)">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';

export default {
  emits: ['edit-book'],
  setup(_, { emit }) {
    const books = ref([]);

    const fetchBooks = async () => {
      const response = await axios.get('/api/books');
      books.value = response.data;
    };

    const deleteBook = async (id) => {
      await axios.delete(`/api/books/${id}`);
      fetchBooks();
    };

    const editBook = (book) => {
      emit('edit-book', book);
    };

    onMounted(fetchBooks);

    return {
      books,
      fetchBooks,
      deleteBook,
      editBook,
    };
  },
};
</script>

<style scoped>
ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
}

button {
  margin-left: 10px;
}
</style>
