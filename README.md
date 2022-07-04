# sliced

The stream slicer (useful for parallel testing) [![Go](https://github.com/udzura/sliced/actions/workflows/go.yml/badge.svg)](https://github.com/udzura/sliced/actions/workflows/go.ym)

## Install

```
go get github.com/udzura/sliced
```

## Usage

```
$ find test -name 'test_*.rb' | sort | head -20
test/-ext-/arith_seq/test_arith_seq_extract.rb
test/-ext-/array/test_resize.rb
test/-ext-/array/test_to_ary_concat.rb
test/-ext-/bignum/test_big2str.rb
test/-ext-/bignum/test_bigzero.rb
test/-ext-/bignum/test_div.rb
test/-ext-/bignum/test_mul.rb
test/-ext-/bignum/test_pack.rb
test/-ext-/bignum/test_str2big.rb
test/-ext-/bug_reporter/test_bug_reporter.rb
test/-ext-/class/test_class2name.rb
test/-ext-/debug/test_debug.rb
test/-ext-/debug/test_profile_frames.rb
test/-ext-/exception/test_data_error.rb
test/-ext-/exception/test_enc_raise.rb
test/-ext-/exception/test_ensured.rb
test/-ext-/exception/test_exception_at_throwing.rb
test/-ext-/file/test_stat.rb
test/-ext-/float/test_nextafter.rb

$ find test -name 'test_*.rb' | sort | wc -l
     837
```

Then:

```
$ find test -name 'test_*.rb' | sort | sliced -n 20 -i 0 --seed 123
2022/07/04 23:32:10 start: %!d(float64=0) end:%!d(float64=42) 
test/openssl/test_pkey.rb
test/rdoc/test_rdoc_parser_changelog.rb
test/win32ole/test_win32ole_variant_outarg.rb
test/ruby/enc/test_case_options.rb
test/rubygems/test_gem_util.rb
test/optparse/test_summary.rb
test/ruby/test_condition.rb
test/ruby/test_exception.rb
test/logger/test_formatter.rb
test/racc/test_grammar_file_parser.rb
...


$ find test -name 'test_*.rb' | sort | sliced -n 20 -i 10 --seed 123 | head 
2022/07/04 23:32:42 start: %!d(float64=419) end:%!d(float64=461) 
test/mkmf/test_convertible.rb
test/ruby/test_super.rb
test/-ext-/symbol/test_inadvertent_creation.rb
test/win32ole/test_win32ole_record.rb
test/date/test_date_arith.rb
test/mkmf/test_mkmf.rb
test/rdoc/test_rdoc_markup_to_label.rb
test/rdoc/test_rdoc_markup_parser.rb
test/ruby/enc/test_gbk.rb
test/rdoc/test_rdoc_rubygems_hook.rb
...

$ find test -name 'test_*.rb' | sort | sliced -n 20 -i 0 --seed 123 | wc -l
2022/07/04 23:32:27 start: %!d(float64=0) end:%!d(float64=42) 
      42
	  
$ find test -name 'test_*.rb' | sort | sliced -n 20 -i 10 --seed 123 | wc -l
2022/07/04 23:32:55 start: %!d(float64=419) end:%!d(float64=461) 
      42
```

## Reference

* [Test splitting and parallelism - CircleCI docs](https://circleci.com/docs/2.0/parallelism-faster-jobs)

## License

MIT
