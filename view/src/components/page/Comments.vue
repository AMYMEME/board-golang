<template>
    <div>
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
                    <b-button v-b-modal.update_comment class="mr-1" variant="outline-primary" cols="auto" @click="passContents(comment.id, comment.contents)">수정</b-button>
                    <b-button variant="outline-danger" cols="auto" @click="deleteComment(comment.id)">삭제</b-button>
                </b-list-group-item>
            </b-list-group>
        </b-card>

        <b-modal centered id="update_comment" title="댓글 수정" @ok="updateComment(update.comment_id)">
            <b-form-input v-bind:value="update.contents" v-model="update.contents">
            </b-form-input>
        </b-modal>
    </div>
</template>

<script>
    export default {
        name: "Comments",
        data() {
            return {
                comments: [],
                form: {
                    member_id: null,
                    contents: '',
                },
                update: {
                    comment_id: null,
                    contents: '',
                },
                error: '',
            }
        },
        created() {
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
            updateComment() {
                this.$axios.patch(`http://localhost:8090/comment/${this.update.comment_id}`, {
                    contents: this.update.contents
                })
                    .then(() => {
                        location.reload();
                    })
                    .catch((err) => {
                        this.error = JSON.stringify(err);
                    });
            },
            passContents(commentID, commentContents) {
                this.update.contents = commentContents;
                this.update.comment_id = commentID;
            },
        }
    }
</script>

<style scoped>

</style>