require "erb"

tag = ENV.fetch("RELEASE_TAG")

unless tag.match?(/\Av\d+\.\d+\.\d+\z/)
  abort "Expected a release tag such as v0.2.0"
end

version = tag.delete_prefix("v")

checksums = {}

expected_files = [
  "portman_darwin_arm64.tar.gz",
  "portman_darwin_amd64.tar.gz"
]

File.foreach("release-assets/checksums.txt") do |line|
  fields = line.split
  abort "Malformed checksum entry" unless fields.length == 2

  checksum, filename = fields

  unless expected_files.include?(filename)
    abort "Unexpected archive: #{filename}"
  end

  if checksums.key?(filename)
    abort "Duplicate checksum entry: #{filename}"
  end

  checksums[filename] = checksum
end

arm64_sha = checksums.fetch("portman_darwin_arm64.tar.gz")
amd64_sha = checksums.fetch("portman_darwin_amd64.tar.gz")

[arm64_sha, amd64_sha].each do |checksum|
  unless checksum.match?(/\A[0-9a-fA-F]{64}\z/)
    abort "Expected a 64-character SHA-256 checksum"
  end
end

template = ERB.new(File.read("packaging/portman.rb.erb"))

formula = template.result_with_hash(
  version: version,
  tag: tag,
  arm64_sha: arm64_sha,
  amd64_sha: amd64_sha
)

puts formula