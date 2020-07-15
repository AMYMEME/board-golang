<template>
  <div>
    <h3 class="mb-3">POST LIST</h3>
    <b-container fluid>
      <div class="col-12">

        <b-card no-body>
          <b-table striped hover @row-clicked="myRowClickHandler" id="post-table"
                   :items="posts" :fields="fields"
                   :per-page="perPage" :current-page="currentPage">
            <!-- A custom formatted column -->
            <template v-slot:cell(datetime)="data">
              <a>{{ formatDatetime(data.value) }}</a>
            </template>
          </b-table>
          <b-pagination
                  v-model="currentPage"
                  :total-rows="this.posts.length"
                  :per-page="perPage"
                  aria-controls="post-table">
          </b-pagination>
        </b-card>
      </div>
    </b-container>
    <b-button class="mt-3" href="/new/post">글쓰기</b-button>
  </div>
</template>

<script>
  export default {
    name: 'get-all-posts',
    data() {
      return {
        perPage: 3,
        currentPage: 1,
        fields: [
          {
            key: 'member_id',
            label: '작성자',
            sortable: false
          },
          {
            key: 'title',
            label: '제목',
            sortable: false
          },
          {
            key: 'datetime',
            label: '일시',
            sortable: true,
          }
        ],
        posts: [],
        error: ''
      }
    },
    created() {
      this.$axios.get('http://localhost:8090/posts')
              .then((res) => {
                this.posts = res.data;
              })
              .catch((err) => {
                this.error = JSON.stringify(err);
              });
    },
    methods: {
      formatDatetime(datetime) {
        return datetime.slice(0, 10) + ' ' + datetime.slice(11, -4)
      },
      myRowClickHandler(record) {
        location.href = '/post/'+record.id
      }
    }
  }
</script>

<style>

</style>