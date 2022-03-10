# frozen_string_literal: true

require 'httparty'
require 'nokogiri'
require 'csv'
require 'smarter_csv'

csvs = SmarterCSV.process("./member.csv")
csvs.each do |csv|
  link = csv[:profile_link]
  response = HTTParty.get(link)
  body = response.body
  html_doc = Nokogiri::HTML5(body)
  image_link = html_doc.css('.profile_picts img').first['src']

  profile = html_doc.css('.profile_detail').first

  name = profile.css('.name_j').text
  p profile.css("dl dt").count
end
