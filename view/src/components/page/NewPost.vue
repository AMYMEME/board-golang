<template>
  <div>
    <b-container fluid>
      <div class="col-12">
        <b-card header="NEW POST">
        <b-form @submit="newPost">
          <b-form-input placeholder="title" id="title" v-model="form.title" required></b-form-input>
          <b-form-textarea placeholder="contents" class="mt-3" id="contents" v-model="form.contents" rows="16"required></b-form-textarea>
          <b-button class="mt-3 mr-1" href="/posts" variant="outline-secondary">취소</b-button>
          <b-button class="mt-3" type="submit" variant="primary">게시물 등록</b-button>
        </b-form>
        </b-card>
      </div>
    </b-container>
  </div>
</template>
<script>
  export default {
    name: 'new-post',
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
    methods: {
      newPost() {
        this.$axios.post('http://localhost:8090/post', {
          member_id: 2, //TODO: 인증 구현 후 수정 필요
          title: this.form.title,
          contents: this.form.contents
        })
                .then((res)=> {
                  window.location.replace(res.data)
                  alert('새 게시글을 등록했습니다')
                })
                .catch((err) => {
                  this.error = JSON.stringify(err)
                  alert('에러가 발생했습니다 : '+this.error)
                })
      },
    }
  }
</script>

<style>

</style>