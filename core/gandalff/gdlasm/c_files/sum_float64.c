#include <immintrin.h>

void sum_float64_avx_intrinsics(double *vec, int *size, double *res)
{
    int rem = (*size) % 4;

    __m256d acc = _mm256_set1_pd(0);
    for (int i = 0; i < (*size) - rem; i += 4)
    {
        __m256d v = _mm256_load_pd(&vec[i]);
        acc = _mm256_add_pd(acc, v);
    }

    acc = _mm256_hadd_pd(acc, acc); // a[0] = a[0] + a[1], a[2] = a[2] + a[3]
    *res = _mm256_cvtsd_f64(acc) + _mm_cvtsd_f64(_mm256_extractf128_pd(acc, 1));

    // Tail
    for (int i = (*size) - rem; i < (*size); i++)
    {
        *res += vec[i];
    }
}

void sum_float64_c(double *vec, int *size, double *res)
{
    for (int i = 0; i < *size; i++)
    {
        *res += vec[i];
    }
}

void sum_grouped_float64_c(double *vec, int *vecSize, int *indeces, int *indecesSize, double *res)
{
    for (int i = 0; i < *indecesSize; i++)
    {
        int index = indeces[i];
        *res += vec[index];
    }
}


// void sum_float64_grouped_avx_intrinsics(double vec[], size_t len, double *res)
// {
//     __m256d acc = _mm256_set1_pd(0);
//     for (int i = 0; i < len; i += 4)
//     {
//         __m256d v = _mm256_load_pd(&vec[i]);
//         acc = _mm256_add_pd(acc, v);
//     }

//     acc = _mm256_hadd_pd(acc, acc); // a[0] = a[0] + a[1], a[2] = a[2] + a[3]
//     *res = _mm256_cvtsd_f64(acc) + _mm_cvtsd_f64(_mm256_extractf128_pd(acc, 1));
// }