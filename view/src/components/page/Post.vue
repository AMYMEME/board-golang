<template>
  <div>
    <b-container fluid>
      <div class="col-12">
        <b-card no-body>
          <b-card-header header-tag="nav">
            <b-nav card-header pills>
              <b-nav-item disabled>작성자 : </b-nav-item>
              <b-nav-item>{{writer}}</b-nav-item>
              <b-nav-item class="mr-auto" disabled>일시 : {{formatDatetime(datetime)}}</b-nav-item>
              <b-button class="mr-1" variant="outline-primary" cols="auto" @click="">수정</b-button>
              <b-button variant="outline-danger" cols="auto" @click="deletePost">삭제</b-button>
            </b-nav>
          </b-card-header>
          <b-card-body class="text-center">
            <b-card-title>{{title}}</b-card-title>
            <b-card-text>{{contents}}</b-card-text>
          </b-card-body>
        </b-card>
        <b-row class="mt-3">
          <b-button cols="auto" class="mr-auto" href="/posts" variant="outline-secondary">목록으로</b-button>
        </b-row>
      </div>

      <div class="col-12">
        <b-card no-body class="mt-3">
          <b-card-header>
            <b-form @submit="newComment">
              <b-form-input placeholder="새 댓글 내용" v-model="form.contents" required></b-form-input>
              <b-button type="submit" class="mt-2 ml-auto">댓글 등록</b-button>
            </b-form>
          </b-card-header>
          <b-card-body v-if="!comments">
            등록된 댓글이 없습니다.
          </b-card-body>
          <b-list-group flush>
            <b-list-group-item v-for="comment in comments" class="d-flex align-items-center">
              <b-avatar class="mr-3"></b-avatar>
              <a href="#" type="button">{{comment.member_id}}</a>
              <span class="ml-2 mr-auto">{{comment.contents}}</span>
              <span class="mr-2 ">{{formatDatetime(comment.datetime)}}</span>
              <b-button class="mr-1" variant="outline-primary" cols="auto" @click="">수정</b-button>
              <b-button variant="outline-danger" cols="auto" @click="deleteComment(comment.id)">삭제</b-button>
            </b-list-group-item>
          </b-list-group>
        </b-card>

      </div>


    </b-container>
  </div>
</template>
<script>
  export default {
    name: 'get-post',
    data() {
      return {
        visible: true,
        writer: '',
        title: '',
        contents: '',
        datetime: '',
        comments: [],
        form:{
          member_id: null,
          contents: '',
        },
        error: ''
      }
    },
    created() {
      this.$axios.get(`http://localhost:8090/post/${this.$route.params.id}`)
              .then((res) => {
                this.writer = res.data.member_id;
                this.title = res.data.title;
                this.contents = res.data.contents;
                this.datetime = res.data.datetime;
                this.writer = res.data.member_id;
              })
              .catch((err) => {
                this.error = JSON.stringify(err);
              });

      this.$axios.get(`http://localhost:8090/post/${this.$route.params.id}/comments`)
              .then((res) => {
                this.comments = res.data;
              })
              .catch((err) => {
                this.error = JSON.stringify(err);
              });
    },
    methods: {
      formatDatetime(datetime) {
        return datetime.slice(0, 10) + ' ' + datetime.slice(11, -4)
      },
      deletePost() {
        this.$axios.delete(`http://localhost:8090/post/${this.$route.params.id}`)
                .then((res)=> {
                  location.replace('/posts');
                })
                .catch((err) => {
                  this.error = JSON.stringify(err);
                });
      },
      newComment() {
        this.$axios.post(`http://localhost:8090/post/${this.$route.params.id}/comment`, {
          member_id: 2, //TODO: 인증 구현 후 수정 필요
          contents: this.form.contents
        })
                .then(() => {
                  location.reload();
                })
                .catch((err) => {
                  this.error = JSON.stringify(err);
                });
      },
      deleteComment(commentID) {
        this.$axios.delete(`http://localhost:8090/comment/${commentID}`)
                .then(() => {
                  location.reload();
                })
                .catch((err) => {
                  this.error = JSON.stringify(err);
                });
      },
    }
  }
</script>

<style>

</style>