resource "local_file" "file" {
  content  = "test"
  filename = "${path.module}/test.txt"
}