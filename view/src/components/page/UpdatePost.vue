<template>
  <div>
    <b-container fluid>
      <div class="col-12">
        <b-card header="EDIT POST">
        <b-form @submit="updatePost">
          <b-form-input placeholder="title" id="title" value="form.title" v-model="form.title" required></b-form-input>
          <b-form-textarea placeholder="contents" class="mt-3" value="form.contents" id="contents" v-model="form.contents" rows="16"required></b-form-textarea>
          <b-button class="mt-3" type="submit" variant="primary">게시물 수정</b-button>
        </b-form>
        </b-card>
      </div>
    </b-container>
  </div>
</template>
<script>
  export default {
    name: 'update-post',
    data() {
      return {
        form:{
          member_id: null,
          title: '',
          contents: '',
        },
        error: ''
      }
    },
    created() {
      this.$axios.get(`http://localhost:8090/post/${this.$route.params.id}`)
              .then((res) => {
                this.form.title = res.data.title;
                this.form.contents = res.data.contents;
              })
              .catch((err) => {
                this.error = JSON.stringify(err);
              });
    },
    methods: {
      updatePost() {
        this.$axios.put(`http://localhost:8090/post/${this.$route.params.id}`, {
          member_id: 2, //TODO: 인증 구현 후 수정 필요
          title: this.form.title,
          contents: this.form.contents
        })
                .then((res) =>{
                  this.$router.push(res.data)
                })
                .catch((err) => {
                  this.error = JSON.stringify(err);
                })
                .finally((res) => {
                  alert('수정 완료');
                });
      },
    }
  }
</script>

<style>

</style>