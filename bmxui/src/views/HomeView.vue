<template>
  <div>
    <BookForm @create-book="createBook" @update-book="updateBook" :bookToEdit="bookToEdit" @clear-edit="clearEdit" />
    <BookList @edit-book="setBookToEdit" />
  </div>
</template>

<script>
import { ref } from 'vue';
import BookForm from '../components/BookForm.vue';
import BookList from '../components/BookList.vue';
import axios from 'axios';

export default {
  components: {
    BookForm,
    BookList,
  },
  setup() {
    const bookToEdit = ref(null);

    const createBook = async (book) => {
      await axios.post('/api/books', book);
      clearEdit();
    };

    const updateBook = async (book) => {
      await axios.put(`/api/books/${book.id}`, book);
      clearEdit();
    };

    const setBookToEdit = (book) => {
      bookToEdit.value = book;
    };

    const clearEdit = () => {
      bookToEdit.value = null;
    };

    return {
      bookToEdit,
      createBook,
      updateBook,
      setBookToEdit,
      clearEdit,
    };
  },
};
</script>
