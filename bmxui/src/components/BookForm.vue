<template>
  <div>
    <h2>{{ isEdit ? 'Edit Book' : 'Add New Book' }}</h2>
    <form @submit.prevent="handleSubmit">
      <div>
        <label for="title">Title:</label>
        <input v-model="book.title" id="title" required />
      </div>
      <div>
        <label for="author">Author:</label>
        <input v-model="book.author" id="author" required />
      </div>
      <button type="submit">{{ isEdit ? 'Update' : 'Create' }}</button>
      <button type="button" @click="clearForm">Cancel</button>
    </form>
  </div>
</template>

<script>
import { ref, watch } from 'vue';

export default {
  props: ['bookToEdit'],
  setup(props, { emit }) {
    const book = ref({ title: '', author: '' });
    const isEdit = ref(false);

    watch(
      () => props.bookToEdit,
      (newVal) => {
        if (newVal) {
          book.value = { ...newVal };
          isEdit.value = true;
        } else {
          clearForm();
        }
      }
    );

    const handleSubmit = () => {
      if (isEdit.value) {
        emit('update-book', book.value);
      } else {
        emit('create-book', book.value);
      }
      clearForm();
    };

    const clearForm = () => {
      book.value = { title: '', author: '' };
      isEdit.value = false;
      emit('clear-edit');
    };

    return {
      book,
      isEdit,
      handleSubmit,
      clearForm,
    };
  },
};
</script>

<style scoped>
form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

button {
  margin-top: 10px;
}
</style>
