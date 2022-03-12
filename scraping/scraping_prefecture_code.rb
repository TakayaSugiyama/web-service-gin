require 'httparty'
require 'nokogiri'
require 'csv'

res = HTTParty.get('https://nlftp.mlit.go.jp/ksj/gml/codelist/PrefCd.html')
body = res.body

dom = Nokogiri::HTML5(body)
prefectures = dom.css("tr")[1..(dom.count-1)].map do |t|
  t.text.strip.split("\n").map(&:strip)
end

data = prefectures.flatten.map do |p|
  p.scan(/(\d+)(.+)/)
end

CSV.open("prefectures.csv", 'w') do |csv|
  data.sort.map(&:flatten).each do |d|
    next if d.empty?
    csv << [d[0].to_i, d[1]]
  end
end



