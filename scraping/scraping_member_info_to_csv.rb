# frozen_string_literal: true

require 'httparty'
require 'nokogiri'
require 'csv'
require 'smarter_csv'

result = []
csvs = SmarterCSV.process("./member.csv")
csvs.each do |csv|
  link = csv[:profile_link]
  response = HTTParty.get(link)
  body = response.body
  html_doc = Nokogiri::HTML5(body)
  image_link = html_doc.css('.profile_picts img').first['src']

  profile = html_doc.css('.profile_detail').first

  data = {}

  name = profile.css('.name_j').text
  data[:name] = name

  profile_doc = profile.css("dl dd")

  nickname = profile_doc[0].text
  data[:nickname] = nickname
  data[:image_link] = image_link

  birthday = profile_doc[1].text.scan(/(\d+)年(\d+)月(\d+)日/).flatten.map(&:to_i)
  data[:birthday_year] = birthday[0]
  data[:birthday_month] = birthday[1]
  data[:birthday_day] = birthday[2]

  data[:blood_type] = profile_doc[2].text
  data[:place_of_birth] = profile_doc[3].text
  data[:height] = profile_doc[4].text.scan(/\d+/).first.to_i
  data[:hobby] = profile_doc[5].text
  data[:special_skill] = profile_doc[6].text
  data[:best_feature] = profile_doc[7].text
  result << data
end

CSV.open("member_info.csv", "w") do |csv|
  csv << result.first.keys
  result.each do |r|
    csv << r.values
  end
end

