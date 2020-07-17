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
              <b-button class="mr-1" variant="outline-primary" cols="auto" :href="`/post/${this.$route.params.id}/update`">수정</b-button>
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
        <comments>

        </comments>
      </div>
    </b-container>
  </div>
</template>
<script>
  import Comments from "./Comments";
  export default {
    name: 'get-post',
    data() {
      return {
        writer: '',
        title: '',
        contents: '',
        datetime: '',
        error: ''
      }
    },
    components: {
      Comments
    },
    created() {
      this.$axios.get(`http://localhost:8090/post/${this.$route.params.id}`)
              .then((res) => {
                this.writer = res.data.member_id;
                this.title = res.data.title;
                this.contents = res.data.contents;
                this.datetime = res.data.datetime;
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
      }
    }
  }
</script>

<style>

</style>